export interface ObjectNode {
  id: string;
  name: string;
  fullPath: string;
  isFolder: boolean;
  sizeBytes: number;
  objectCount: number;
  lastModified?: string;
  children: ObjectNode[];
}

export function buildObjectTree(
  objects: Array<{ key?: string; bucketName?: string; name?: string; totalSizeBytes?: number; sizeBytes?: number; lastModified?: string; creationDate?: string }>
): ObjectNode[] {
  const root: ObjectNode[] = [];

  objects.forEach((obj: any, idx) => {
    const rawKey = obj.key || obj.name || obj.bucketName || `object-${idx + 1}`;
    const size = obj.totalSizeBytes ?? obj.sizeBytes ?? obj.size ?? 0;
    const modified = obj.lastModified || obj.creationDate || '';

    const parts = rawKey.split('/').filter(Boolean);
    if (parts.length === 0) return;

    let currentLevel = root;
    let currentPath = '';

    parts.forEach((part: string, index: number) => {
      const isLast = index === parts.length - 1;
      const isFolder = !isLast || rawKey.endsWith('/');
      currentPath = currentPath ? `${currentPath}/${part}` : part;

      let existingNode = currentLevel.find((n) => n.name === part && n.isFolder === isFolder);

      if (!existingNode) {
        existingNode = {
          id: `${currentPath}-${isFolder ? 'folder' : 'file'}-${idx}`,
          name: part,
          fullPath: currentPath,
          isFolder,
          sizeBytes: isFolder ? 0 : size,
          objectCount: isFolder ? 0 : 1,
          lastModified: modified,
          children: []
        };
        currentLevel.push(existingNode);
      }

      if (isFolder) {
        existingNode.sizeBytes += size;
        existingNode.objectCount += 1;
        currentLevel = existingNode.children;
      }
    });
  });

  // Sort nodes: folders first, then files alphabetically
  const sortNodes = (nodes: ObjectNode[]) => {
    nodes.sort((a, b) => {
      if (a.isFolder && !b.isFolder) return -1;
      if (!a.isFolder && b.isFolder) return 1;
      return a.name.localeCompare(b.name);
    });
    nodes.forEach((n) => {
      if (n.children.length > 0) sortNodes(n.children);
    });
  };

  sortNodes(root);
  return root;
}

export function formatBytes(bytes: number, decimals: number = 2): string {
  if (bytes === 0) return '0 Bytes';
  const k = 1024;
  const dm = decimals < 0 ? 0 : decimals;
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`;
}
