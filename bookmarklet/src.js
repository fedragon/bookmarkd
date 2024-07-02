const tags = prompt("tags: ", "")
  .split(",")
  .map(tag => 'tags=' + encodeURIComponent(tag.trim()))
  .join("&");

const addr = 'http://localhost:3333/api/bookmarks';
const vault = 'my-vault';
const folder = 'Clippings';

let url = addr + '?' +
  'vault=' + encodeURIComponent(vault) +
  '&folder=' + encodeURIComponent(folder) +
  '&url=' + encodeURIComponent(document.location);

if (tags) {
  url += '&' + tags
}

document.location.href = url;
