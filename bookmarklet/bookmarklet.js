javascript:(function()%7Bconst%20tags%20%3D%20prompt(%22tags%3A%20%22%2C%20%22%22)%0A%20%20.split(%22%2C%22)%0A%20%20.map(tag%20%3D%3E%20'tags%3D'%20%2B%20encodeURIComponent(tag.trim()))%0A%20%20.join(%22%26%22)%3B%0A%0Aconst%20addr%20%3D%20'http%3A%2F%2Flocalhost%3A3333%2Fapi%2Fbookmarks'%3B%0Aconst%20vault%20%3D%20'my-vault'%3B%0Aconst%20folder%20%3D%20'Clippings'%3B%0A%0Alet%20url%20%3D%20addr%20%2B%20'%3F'%20%2B%0A%20%20'vault%3D'%20%2B%20encodeURIComponent(vault)%20%2B%0A%20%20'%26folder%3D'%20%2B%20encodeURIComponent(folder)%20%2B%0A%20%20'%26url%3D'%20%2B%20encodeURIComponent(document.location)%3B%0A%0Aif%20(tags)%20%7B%0A%20%20url%20%2B%3D%20'%26'%20%2B%20tags%0A%7D%0A%0Adocument.location.href%20%3D%20url%3B%7D)()%3B