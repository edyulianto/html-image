package main

import (
	"context"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	// Path ke file HTML
	htmlFilePath := "id-card-design-ktp.html"

	// Baca isi file HTML
	htmlData, err := os.ReadFile(htmlFilePath)
	if err != nil {
		log.Fatalf("Failed to read HTML file: %v", err)
	}

	// Buat context Chrome
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Gunakan timeout untuk mencegah infinite wait
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := chromedp.Run(ctx,
		// the navigation will trigger the "page.EventLoadEventFired" event too,
		// so we should add the listener after the navigation.
		chromedp.Navigate("about:blank"),
		// set the page content and wait until the page is loaded (including its resources).
		chromedp.ActionFunc(func(ctx context.Context) error {
			lctx, cancel := context.WithCancel(ctx)
			defer cancel()
			var wg sync.WaitGroup
			wg.Add(1)
			chromedp.ListenTarget(lctx, func(ev interface{}) {
				if _, ok := ev.(*page.EventLoadEventFired); ok {
					// It's a good habit to remove the event listener if we don't need it anymore.
					cancel()
					wg.Done()
				}
			})

			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}

			if err := page.SetDocumentContent(frameTree.Frame.ID, string(htmlData)).Do(ctx); err != nil {
				return err
			}
			wg.Wait()
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(true).Do(ctx)
			if err != nil {
				return err
			}
			return ioutil.WriteFile("sample.pdf", buf, 0644)
		}),
	); err != nil {
		log.Fatal(err)
	}
}
