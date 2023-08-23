/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: false,
    transpilePackages: ["@mui/material"],
    experimental: {
        appDir: true,
    },
    compiler: {
        styledComponents: {
            displayName: false,
            ssr: true
        },
    },
}

module.exports = nextConfig
