package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

type Resource struct {
	Filename string
	Url      string
}

type Downloader struct {
	wg         *sync.WaitGroup
	pool       chan *Resource
	Concurrent int
	HttpClient http.Client
	TargetDir  string
	Resources  []Resource
}

func NewDownloader(targetDir string) *Downloader {
	concurrent := runtime.NumCPU()
	return &Downloader{
		wg:         &sync.WaitGroup{},
		TargetDir:  targetDir,
		Concurrent: concurrent,
	}
}

func (d *Downloader) AppendResource(filename, url string) {
	d.Resources = append(d.Resources, Resource{
		Filename: filename,
		Url:      url,
	})
}

func (d *Downloader) Download(resource Resource, progress *mpb.Progress) error {
	defer d.wg.Done()
	d.pool <- &resource
	finalPath := d.TargetDir + "/" + resource.Filename

	target, err := os.Create(finalPath + ".tmp")
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodGet, resource.Url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		target.Close()
		return err
	}
	defer resp.Body.Close()
	fileSize, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	bar := progress.Add(
		int64(fileSize),
		mpb.NewBarFiller(mpb.BarStyle().Lbound("[").Filler("=").Tip(">").Padding("░").Rbound("]")),
		mpb.PrependDecorators(
			decor.Name(resource.Filename, decor.WC{W: len(resource.Filename) + 1, C: decor.DidentRight}),

			decor.OnComplete(
				decor.AverageETA(decor.ET_STYLE_GO, decor.WC{W: 4}), " СКАЧАН ",
			),
		),

		mpb.AppendDecorators(
			decor.CountersKibiByte("% .2f / % .2f   "),
		),
	)
	reader := bar.ProxyReader(resp.Body)
	defer reader.Close()
	if _, err := io.Copy(target, reader); err != nil {
		target.Close()
		return err
	}

	target.Close()
	if err := os.Rename(finalPath+".tmp", finalPath); err != nil {
		return err
	}

	<-d.pool
	return nil
}

func (d *Downloader) Start() error {
	d.pool = make(chan *Resource, d.Concurrent)
	p := mpb.New(mpb.WithWaitGroup(d.wg))
	for _, resource := range d.Resources {
		d.wg.Add(1)
		go d.Download(resource, p)
	}
	p.Wait()
	d.wg.Wait()

	return nil
}

func Unzip(src string, dest string) ([]string, error) {
	var filenames []string
	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()
	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}
		filenames = append(filenames, fpath)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}
		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
