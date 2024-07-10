const tags = prompt("tags: ", "")
  .split(",")
  .map(tag => 'tag=' + encodeURIComponent(tag.trim()))
  .join("&");

const addr = '{{.Address}}';
const vault = '{{.Vault}}';
const folder = '{{.Folder}}';

let url = addr + '?' +
  'vault=' + encodeURIComponent(vault) +
  '&folder=' + encodeURIComponent(folder) +
  '&url=' + encodeURIComponent(document.location);

if (tags) {
  url += '&' + tags
}

document.location.href = url;

