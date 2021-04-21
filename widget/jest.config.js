module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'jsdom',
  transform: {
    '^.+\\.svelte$': [
      'svelte-jester',
      {
        preprocess: true,
      },
    ],
    '^.+\\.ts$': 'ts-jest',
  },
  moduleNameMapper: {
    '\\.svg': '<rootDir>/__mocks__/svgrMock.js',
  },
  moduleFileExtensions: ['js', 'ts', 'svelte'],
  setupFilesAfterEnv: ['<rootDir>setupTests.js'],
}
