import {chromium} from 'playwright';
import * as nhp from 'node-html-parser';
import * as fs from "node:fs";

(async () => {
    const browser = await chromium.connectOverCDP('http://localhost:1243');
    const defaultContext = browser.contexts()[0];
    const page = defaultContext.pages()[0];

    const serverAddress = 'http://localhost:11235';
    const vault = 'my-vault';
    const folder = 'Clippings';
    const doc = fs.readFileSync('./pocket-export.html', 'utf8');
    const root: nhp.HTMLElement = nhp.parse(doc);

    const links = root.querySelectorAll('a').map(
        (a: nhp.HTMLElement): { url: string, link: string } => {
            const tags = a.getAttribute('tags')
                .split(",")
                .map(tag => 'tag=' + encodeURIComponent(tag.trim()))
                .join("&");

            const params = [
                `url=${encodeURIComponent(a.getAttribute('href'))}`,
                `vault=${encodeURIComponent(vault)}`,
                `folder=${encodeURIComponent(folder)}`,
                `epoch=${a.getAttribute('time_added')}`,
                'silent=true',
                tags
            ].join('&')

            return {url: a.getAttribute('href'), link: `${serverAddress}/api/bookmarks?${params}`}
        }
    );

    for (const {url, link} of links) {
        try {
            await page.goto(link);
        } catch (err) {
            if (err instanceof Error && !err.message.startsWith('page.goto: net::ERR_ABORTED')) {
                console.log('failed to open', url, err.name);
            }
        }
    }
    await browser.close();
})();
