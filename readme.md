# goscrape

An AWS lambda-based web scraper - when called, returns JSON about all attributes of a certain DOM element.

Built after I wanted to save images from [Visvim](http://visvim.tv) for inspiration but couldn't as image downloading was blocked.

### Example
```
curl -XPOST -d "url=https://www.visvim.tv/lookbook/ss18/visvim/" \
	-d "element=img" \
	-d "attribute=href" \
	https://example-goscraper-gateway.com/scrape
```

### Use Cases:
- Getting image hrefs for download
- Getting script srcs for logging/download
- Getting a list of all classes/ids on a certain element

