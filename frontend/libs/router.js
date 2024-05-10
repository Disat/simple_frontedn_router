export default class Router {

    constructor(routeMap) {
        this.routeMap = routeMap;
        this._bindPopState();
    }

    init(path) {
        // path = Router.correctPath(path);
        history.replaceState({ path: path }, '', path);
        this.routeMap[path] && this.routeMap[path]();
    }

    go(path) {

        // path = Router.correctPath(path);
        history.pushState({ path: path }, '', path);
        this.routeMap[path] && this.routeMap[path]();
    }

    _bindPopState() {
        window.addEventListener('popstate', (e) => {
            const path = e.state && e.state.path;
            this.routeMap[path] && this.routeMap[path]();
        });
    }

    // static correctPath(path) {
    //     if (path !== '/' && path.slice(-1) === '/') {
    //         path = path.match(/(.+)\/$/)[1];
    //     }
    //     return path;
    // }
}
