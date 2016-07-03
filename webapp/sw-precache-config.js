module.exports = {
    staticFileGlobs: [
        '/index.html',
        '/manifest.json',
        '/static/bower_components/webcomponentsjs/webcomponents-lite.min.js',
        '/static/src/**/*'
    ],
    handleFetch: true,
    navigateFallback: '/index.html',
    runtimeCaching: [
        {
        urlPattern: /\/api\//,
        handler: 'networkOnly'
        },
        {
            urlPattern: /\//,
            handler: 'fastest',
            options: {
                name: "site-cache"
            }
        }
    ]
};
