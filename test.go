package main

// func getImageSizes(url string) (map[string]string, error) {
// 	imageSizes := make(map[string]string)

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	doc, err := html.Parse(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var extractImageSizes func(*html.Node)
// 	extractImageSizes = func(n *html.Node) {
// 		if n.Type == html.ElementNode && n.Data == "img" {
// 			for _, attr := range n.Attr {
// 				if attr.Key == "src" {
// 					imageURL := attr.Val
// 					// Retrieve image size
// 					resp, err := http.Head(imageURL)
// 					if err != nil {
// 						log.Println("Error fetching image:", err)
// 						continue
// 					}
// 					defer resp.Body.Close()
// 					size := resp.Header.Get("Content-Length")
// 					imageSizes[imageURL] = size
// 					break
// 				}
// 			}
// 		}
// 		for c := n.FirstChild; c != nil; c = c.NextSibling {
// 			extractImageSizes(c)
// 		}
// 	}
// 	extractImageSizes(doc)
// 	return imageSizes, nil
// }

// func main() {
// 	url := "https://thengcstore.com/products"
// 	imageSizes, err := getImageSizes(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for imageURL, size := range imageSizes {
// 		fmt.Printf("Image URL: %s, Size: %s\n", imageURL, size)
// 	}
// }
