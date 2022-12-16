/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  images: {
    domains: ["assets.haikara.dev", "localhost"],
  },
};

module.exports = nextConfig;
