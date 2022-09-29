const nextJest = require("next/jest");

const createJestConfig = nextJest({
  // Provide the path to your Next.js app to load next.config.js and .env files in your test environment
  dir: "./",
});

// Add any custom config to be passed to Jest
const customJestConfig = {
  // Add more setup options before each test is run
  // setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
  // if using TypeScript with a baseUrl set to the root directory then you need the below for alias' to work
  moduleDirectories: ["node_modules", "<rootDir>/"],
  testEnvironment: "jest-environment-jsdom",

  setupFilesAfterEnv: ["<rootDir>/jest.setup.js"],
  moduleNameMapper: {
    "^@/components/(.*)$": "<rootDir>/src/components/$1",
    "^@/services/(.*)$": "<rootDir>/src/services/$1",
    "^@/features/(.*)$": "<rootDir>/src/features/$1",
    "^@/app/(.*)$": "<rootDir>/src/app/$1",
    "^@/fixtures/(.*)$": ["<rootDir>/src/fixtures/$1"],
    "^@/lib/(.*)$": ["<rootDir>/src/lib/$1"],
    "^@/pages/(.*)$": ["<rootDir>/src/pages/$1"],
    "^@/styles/(.*)$": ["<rootDir>/src/styles/$1"],
    "^@/types/(.*)$": ["<rootDir>/src/types/$1"],
    "^@/utils/(.*)$": ["<rootDir>/src/utils/$1"],
  },
  // 除外するディレクトリ
  modulePathIgnorePatterns: ["<rootDir>/e2e/"],
};

// createJestConfig is exported this way to ensure that next/jest can load the Next.js config which is async
module.exports = createJestConfig(customJestConfig);
