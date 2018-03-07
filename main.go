// Package main contains the GoScrape lambda function.
package main

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go"
	"log"
	"net/url"
	"string"
)

// Request
type Request struct {
	Url       string `json:"url"`
	Element   string `json:"element"`
	Attribute string `json:attribute"`
}

// Response is the object returned by the scraper, containing the element's name and a slice of attributes
type Response struct {
	Element    string   `json:"element"`
	Attribute  string   `json:"attribute"`
	Attributes []string `json:"attributes"`
}

// initializeDocument takes a url and returns a GoQuery Document struct
func initializeDocument(document_url string) (goquery.Document, error) {
	return goquery.NewDocument(document_url)
}

// findAllElements takes a GoQuery search query and returns a slice of GoQuery Selection structs to act on
func findAllElements(element_name string, document goquery.Document) []goquery.Selection {
	return document.Find(element_name)
}

// getAllAttributes takes a slice of GoQuery Selection structs and plucks out the specified attribute
func getAllAttributes(elements []goquery.Selection, attribute_name string) ([]string, error) {
	all_attributes := make([len(elements)]string)
	elements.Each(func(index int, element *goquery.Selection) {
		attribute, attribute_exists := element.Attr(attribute_name)
		if attribute_exists {
			all_attributes.append(attr)
		}
	})
	return all_attributes
}

// Create some custom errors that can be thrown
var (
	ErrInvalidUrl = errors.New("The URL provided was invalid.")
)

// Scraper is the Lambda function handler
// It calls InitializeDocument, FindAllElements, and then returns JSON with the attribute specified in the request
func Scraper(request Request) (Response, error) {
	response := Response{
		request.Element,
		request.Attribute,
	}
	err := url.ParseRequestURI(request.Url)
	if err != nil {
		panic(err)
	}
	document, err := initializeDocument(document_url)
	if err != nil {
		panic(err)
	}
	elements := findAllElements(request.Element, document)
	attributes := getAllAttributes(elements, request.Attribute)
	response.Attributes = attributes
	return response
}

// main runs the AWS Lambda handler
func main() {
	lambda.Start(Scraper)
}
