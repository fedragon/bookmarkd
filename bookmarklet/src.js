var tags = prompt("tags: ", "")
  .split(",")
  .map(tag => 'tags=' + encodeURIComponent(tag.trim()))
  .join("&");

var addr = 'http://localhost:3333'
var vault = 'my-vault'
var folder = 'Clippings'

var url = addr + '/bookmarks?' +
  'vault=' + encodeURIComponent(vault) +
  '&folder=' + encodeURIComponent(folder) +
  '&url=' + encodeURIComponent(document.location)

if (tags) {
  url += '&' + tags
}

document.location.href = url;
