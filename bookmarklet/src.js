const addr = prompt('HTTP server address: ', 'http://localhost:11235');
const vault = prompt('Vault: ', 'Vault');
const folder = prompt('Folder: ', 'Clippings');

const payload = `
const tags = prompt("tags: ", "")
  .split(",")
  .map(tag => 'tag=' + encodeURIComponent(tag.trim()))
  .join("&");

let url = '${addr}/api/bookmarks?' +
  'vault=' + encodeURIComponent('${vault}') +
  '&folder=' + encodeURIComponent('${folder}') +
  '&url=' + encodeURIComponent(document.location);

if (tags) {
  url += '&' + tags;
}

document.location.href = url;
`

alert("javascript:" + encodeURIComponent("(function(){" + payload + "})();"));
