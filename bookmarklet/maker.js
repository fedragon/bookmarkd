const fs = require('node:fs');

try {
  const code = fs.readFileSync('./src.js', 'utf8');
  content = "javascript:" + encodeURIComponent("(function(){" + code.trim() + "})();");
  fs.writeFileSync('./bookmarklet.js', content);
} catch (err) {
  console.error(err);
}

