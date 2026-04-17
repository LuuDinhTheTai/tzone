export function resolveDeviceImageUrl(imageUrl?: string): string {
  if (!imageUrl) return '';

  // Keep seeded remote URLs unchanged.
  if (/^https?:\/\//i.test(imageUrl)) {
    return imageUrl;
  }

  // Support legacy backslash paths coming from Windows-built data.
  const normalizedPath = imageUrl.replace(/\\/g, '/');

  if (normalizedPath.startsWith('/')) {
    return normalizedPath;
  }

  return `/${normalizedPath}`;
}

