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
                "code": 160,
                "tabWidth": 2,
                "ignoreComments": true,
                "ignoreUrls": true,
                "ignoreStrings": true,
                "ignoreTemplateLiterals": true
            }
        ],
        'space-before-function-paren': 'off',
        'vue/attributes-order': 'off',
        'vue/one-component-per-file': 'off',
        'vue/html-closing-bracket-newline': 'off',
        'vue/max-attributes-per-line': 'off',
        'vue/multiline-html-element-content-newline': 'off',
        'vue/singleline-html-element-content-newline': 'off',
        'vue/attribute-hyphenation': 'off',
        'vue/require-default-prop': 'off',
        'vue/require-explicit-emits': 'off',
        'vue/html-self-closing': [
            'error',
            {
                html: {
                    void: 'always',
                    normal: 'never',
                    component: 'always',
                },
                svg: 'always',
                math: 'always',
            },
        ],
    },
};
