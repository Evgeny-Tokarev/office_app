/** @type {import('next').NextConfig} */
const nextConfig = {
    modularizeImports: {
        '@mui/icons-material/?(((\\w*)?/?)*)': {
            transform: '@mui/icons-material/{{ matches.[1] }}/{{member}}'
        }
    },
    reactStrictMode: false,
    transpilePackages: ["@mui/material"],
    compiler: {
        styledComponents: {
            displayName: false,
            ssr: true
        },
    },
}

module.exports = nextConfig
