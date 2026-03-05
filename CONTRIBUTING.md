# Contributing to VPS Security Control

Thank you for considering contributing! This document provides guidelines and instructions for contributing.

## Code of Conduct

Be respectful, inclusive, and professional in all interactions.

## How to Contribute

### Reporting Bugs

1. Check existing issues to avoid duplicates
2. Create a new issue with:
   - Clear title describing the bug
   - Detailed reproduction steps
   - Expected vs actual behavior
   - System information (OS, Go version, etc.)
   - Screenshots if applicable

### Suggesting Features

1. Check existing issues
2. Create a new issue with:
   - Clear description of the feature
   - Why it's needed (use case)
   - Possible implementation approach
   - Examples or mockups if applicable

### Submitting Pull Requests

1. Fork the repository
2. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. Make your changes:
   - Follow existing code style
   - Add tests for new functionality
   - Update documentation as needed

4. Commit with clear messages:
   ```bash
   git commit -m "Add feature: description of changes"
   ```

5. Push to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```

6. Open a Pull Request with:
   - Clear description of changes
   - Link to related issues
   - Screenshots if UI changes
   - Test results

## Development Setup

### Backend Development

```bash
cd backend
go mod download
cp .env.example .env

# Install air for hot reload
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

### Frontend Development

```bash
cd frontend
npm install
cp .env.example .env
npm run dev
```

## Code Style

### Go
- Follow standard Go conventions
- Run `go fmt` before committing
- Use meaningful variable names
- Add comments for exported functions

### Vue/TypeScript
- Use consistent indentation (2 spaces)
- Use TypeScript for type safety
- Component names in PascalCase
- Use arrow functions where appropriate

## Testing

### Backend Tests
```bash
cd backend
go test -v ./...
```

### Frontend Tests
```bash
cd frontend
npm run test
```

## Documentation

- Update README.md for user-facing changes
- Update API.md for API changes
- Add comments to complex code
- Update INSTALLATION.md for setup changes

## Commit Messages

Use clear, descriptive commit messages:
- `git commit -m "feat: add security event filtering"`
- `git commit -m "fix: resolve database connection pooling issue"`
- `git commit -m "docs: update installation guide"`

Prefixes:
- `feat:` - New feature
- `fix:` - Bug fix
- `docs:` - Documentation
- `style:` - Code style (formatting, missing semicolons, etc.)
- `refactor:` - Code refactoring
- `test:` - Test additions or changes
- `chore:` - Build, dependency updates, etc.

## Pull Request Review

PRs will be reviewed for:
- Code quality and style
- Test coverage
- Documentation updates
- Performance implications
- Security considerations

## Becoming a Maintainer

Active contributors may be invited to become maintainers. Maintainers have:
- Merge permissions
- Release responsibilities
- Code review authority

## Questions?

Feel free to:
- Open an issue for discussion
- Start a discussion in GitHub Discussions (when available)
- Contact the maintainers

---

Thank you for contributing to VPS Security Control! 🚀
