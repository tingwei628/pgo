package internal

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"
)

// default value
const (
	SHA1   = iota // 0
	SHA224        // 1
	SHA256        // 2
	SHA384        // 3
	SHA512        // 4
)
const hashAlgDesc string = "0 (SHA-1)\n1 (SHA-224)\n2 (SHA-256)\n3 (SHA-384)\n4 (SHA-512)\n"
const (
	hashAlgDefault = SHA512
	sizeDefault    = 50
	mutltiple      = 5
	numberHash     = 3
)

type Avatarme struct {
	inputs              string
	hashAlgPtr, sizePtr *int
}

// package init
func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: go run cli.go [options]\n")
		fmt.Fprintf(os.Stderr, "e.g. go run cli.go -hashAlg=2 \n")
		flag.PrintDefaults()
	}
	if len(os.Args) < 2 {
		log.Fatal("expected at least two arguments")
		os.Exit(1)
	}
}
func (opts *Avatarme) SetOpts() {
	flag.StringVar(&opts.inputs, "inputs", "", "inputs desc")
	opts.hashAlgPtr = flag.Int("hashAlg", hashAlgDefault, hashAlgDesc)
	opts.sizePtr = flag.Int("size", sizeDefault, "size desc")
	flag.Parse()
	if *opts.sizePtr < sizeDefault {
		log.Fatalf("size at least 50, but yours is %v", *opts.sizePtr)
	}
}
func (opts *Avatarme) Generate() {
	opts.SetOpts()
	c := [3]uint8{}
	hash := []byte{}
	text := opts.inputs
	hashAlg := *opts.hashAlgPtr
	// get hash
	switch hashAlg {
	case SHA1:
		_a := sha1.Sum([]byte(text))
		hash = _a[:sha1.Size]
	case SHA224:
		_a := sha256.Sum224([]byte(text))
		hash = _a[:sha256.Size224]
	case SHA256:
		_a := sha256.Sum256([]byte(text))
		hash = _a[:sha256.Size]
	case SHA384:
		_a := sha512.Sum384([]byte(text))
		hash = _a[:sha512.Size384]
	case SHA512:
		_a := sha512.Sum512([]byte(text))
		hash = _a[:sha512.Size]
	}
	// select color (first 3 hash value)
	copy(c[:], hash[:3])
	colorrgb := color.RGBA{c[0], c[1], c[2], 0xff} //  R, G, B, Alpha
	// create .png
	outFile, err := os.Create("avatarme.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	// size is mutltiple of 5
	size := int(math.Ceil(float64(*opts.sizePtr)/float64(mutltiple))) * mutltiple
	blockSize := size / mutltiple
	img := image.NewRGBA(image.Rect(0, 0, size, size)) // x1,y1,  x2,y2
	whiteColor := color.RGBA{0xff, 0xff, 0xff, 0xff}   //  R, G, B, Alpha
	// backfill entire surface with white
	draw.Draw(img, img.Bounds(), &image.Uniform{whiteColor}, image.ZP, draw.Src)
	hashLen := len(hash)
	// choose first 25 bytes
	for i := 0; i < (mutltiple * numberHash); i += numberHash {
		hashIndex := i % hashLen
		chunks := []byte{}
		if hashIndex+numberHash < hashLen {
			chunks = append(chunks, hash[hashIndex:hashIndex+3]...)
		} else {
			rest := (hashIndex + numberHash) - hashLen
			chunks = append(chunks, hash[hashIndex:hashLen]...)
			chunks = append(chunks, hash[0:rest]...)
		}
		chunks = append(chunks, chunks[1:2]...)
		chunks = append(chunks, chunks[0:1]...)
		origin := i / numberHash * mutltiple
		for index := origin; index < origin+mutltiple; index += 1 {
			// filter out even number
			if chunks[index-origin]%2 == 1 {
				continue
			}
			horizontal := (index % mutltiple) * blockSize
			vertical := (index / mutltiple) * blockSize
			block := image.Rect(horizontal, vertical, horizontal+blockSize, vertical+blockSize) //  geometry of 2nd rectangle
			draw.Draw(img, block, &image.Uniform{colorrgb}, image.ZP, draw.Src)
		}
	}
	png.Encode(outFile, img)
}
