var { Readability } = require('@mozilla/readability');
var { JSDOM } = require('jsdom');
var url = "https://tonsky.me/blog/good-times-weak-men/";

fetch(url)
  .then(response => {
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    return response.text();
  })
  .then(htmlContent => {
    let doc = new JSDOM(htmlContent, {url: url});
    let reader = new Readability(doc.window.document);
    let article = reader.parse();
    console.log(article.content)
  })
  .catch(error => {
    console.error('There was a problem with the fetch operation:', error);
  });

