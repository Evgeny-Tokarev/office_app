module.exports = {
    env: {
        browser: true,
        es2021: true,
        node: true
    },
    extends: [
        'plugin:@next/next/recommended'
    ],
    parser: '@typescript-eslint/parser',
    parserOptions: {
        ecmaVersion: 'latest',
        sourceType: 'module'
    },
    plugins: [
        '@typescript-eslint',
        'react'
    ],
    rules: {
        'react/prop-types': 'off',
        'react/react-in-jsx-scope': 'off',
        'react/jsx-first-prop-new-line': 'warn',
        'react/jsx-max-props-per-line': ['warn', { when: 'always' }],
        'react/jsx-indent': ['error', 4],
        'semi': ['error', 'never']
    },
    settings: {
        react: {
            version: 'detect'
        }
    }
}
