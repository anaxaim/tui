const path = require('path');

module.exports = {
    env: {
        browser: true,
        es2022: true,
    },
    parserOptions: {
        ecmaVersion: 12,
        sourceType: 'module',
    },
    extends: [
        'plugin:vue/vue3-recommended',
        'airbnb-base',
    ],
    plugins: [
        'vue',
    ],
    settings: {
        'import/resolver': {
            alias: {
                map: [
                    ['@', path.resolve(__dirname, './src')],
                    ['views', path.resolve(__dirname, './src/views')],
                    ['components', path.resolve(__dirname, './src/components')],
                ],
                extensions: ['.js', '.jsx', '.vue', '.json'],
            },
        },
    },
    rules: {
        'vue/multi-word-component-names': 'off',
        "no-console": "off",
        "max-len": [
            "error",
            {
                "code": 120,
                "tabWidth": 2,
                "ignoreComments": true,
                "ignoreUrls": true,
                "ignoreStrings": true,
                "ignoreTemplateLiterals": true
            }
        ]
    },
};
