import Router from "./libs/router.js"

const routeMap = {
    '/': () => {
        const content = document.querySelector('.content');
        content.innerHTML = '<div>welcome to Home Page</div>';
    },
    '/profile': () => {
        const content = document.querySelector('.content');
        content.innerHTML = '<div>welcome to Profile Page</div>';
    },
    '/articles': () => {
        const content = document.querySelector('.content');
        content.innerHTML =
            '<div>' +
            '<p>welcome to Article Page</p>' +
            '<ul>' +
            '<li>文章1</li>' +
            '<li>文章2</li>' +
            '<li>文章3</li>' +
            '</ul>' +
            '</div>';
    }
};

const router = new Router(routeMap);
router.init(location.pathname);
document.querySelector('.menu').addEventListener('click', (e) => {
    if (e.target.tagName === 'A') {
        e.preventDefault();
        router.go(e.target.getAttribute('href'))
    }
});