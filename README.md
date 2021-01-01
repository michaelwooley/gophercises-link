# [Finding links in HTML documents (Gophersices)](https://github.com/gophercises/link)

## Task

Given an HTML document, return the links as a list. 

For example, consider:

```html
<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>! 
      <!-- Comments should not be included -->
    </a>
  </div>
</body>
</html>
```

Should return:

```json
[
    {
        "href": "https://www.twitter.com/joncalhoun",
        "text": "Check me out on twitter"
    }, {
        "href": "https://github.com/gophercises",
        "text": "Gophercises is on Github!"
    }
]
```

## Usage

To run the samples:

```bash
git clone git@github.com:michaelwooley/gophercises-link.git
cd gophercises-link
go build
go run main.go
```
