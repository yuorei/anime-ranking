package application

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/kolesa-team/go-webp/webp"
)

const fileLimitSize = 500000

func UploadGCS(file graphql.Upload) (string, error) {
	webpBuffer, err := ConvertToWebp(file)
	if err != nil {
		return "", err
	}

	bucket := os.Getenv("BUCKET")
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	uu, _ := uuid.NewRandom()
	filename := uu.String() + ".webp"
	o := client.Bucket(bucket).Object(filename)
	// Optional: set a generation-match precondition to avoid potential race
	// conditions and data corruptions. The request to upload is aborted if the
	// object's generation number does not match your precondition.
	// For an object that does not yet exist, set the DoesNotExist precondition.
	o = o.If(storage.Conditions{DoesNotExist: true})
	// If the live object already exists in your bucket, set instead a
	// generation-match precondition using the live object's generation number.
	// attrs, err := o.Attrs(ctx)
	// if err != nil {
	//      return fmt.Errorf("object.Attrs: %v", err)
	// }
	// o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	// Upload an object with storage.Writer.
	wc := o.NewWriter(ctx)
	if _, err = io.Copy(wc, webpBuffer); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	imageURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", os.Getenv("BUCKET"), filename)
	log.Printf("アップロード成功 %s\n", imageURL)
	return imageURL, nil
}

func ConvertToWebp(file graphql.Upload) (*bytes.Buffer, error) {
	if file.Size > fileLimitSize {
		return nil, fmt.Errorf("ファイルサイズが大きすぎます。500KB以下にしてください。")
	}

	var image image.Image
	switch file.ContentType {
	case "image/jpeg":
		// JPEG画像をデコード
		img, err := jpeg.Decode(file.File)
		if err != nil {
			return nil, fmt.Errorf("JPEG画像をデコードに失敗しました")
		}
		image = img

	case "image/png":
		// PNG画像をデコード
		img, err := png.Decode(file.File)
		if err != nil {
			return nil, fmt.Errorf("PNG画像をデコードに失敗しました")
		}
		image = img
	case "image/webp":
		// WEBP画像をデコード
		img, err := webp.Decode(file.File, nil)
		if err != nil {
			return nil, fmt.Errorf("WEBP画像をデコードに失敗しました")
		}
		image = img
	default:
		return nil, fmt.Errorf("対応していないファイルです")
	}

	// WebPに変換
	webpBuffer := new(bytes.Buffer)
	err := webp.Encode(webpBuffer, image, nil)
	if err != nil {
		return nil, fmt.Errorf("WEBP画像にエンコードに失敗しました")
	}

	return webpBuffer, nil
}
