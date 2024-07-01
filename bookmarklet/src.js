var tags = prompt("tags: ", "")
  .split(",")
  .map(tag => 'tags=' + encodeURIComponent(tag.trim()))
  .join("&");

url = 'http://localhost:3000/bookmarks?url=' + encodeURIComponent(document.location)
if (tags) {
  url += '&' + tags
}

document.location.href = url;
