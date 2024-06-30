var tags = prompt("tags: ", "")
  .split(",")
  .map(tag => 'tags=' + tag.trim())
  .join("&");
document.location.href = 'http://localhost:3000/bookmarks?url=' + encodeURIComponent(document.location) + '&' + tags;
