// Deduplication utility for versioned entities
export function deduplicateByName<T extends { name: string; createdAt: string; updatedAt: string }>(
  entities: T[]
): T[] {
  // Group entities by name
  const grouped = entities.reduce((acc, entity) => {
    const name = entity.name;
    if (!acc[name]) {
      acc[name] = [];
    }
    acc[name].push(entity);
    return acc;
  }, {} as Record<string, T[]>);

  // For each group, keep only the latest version
  return Object.values(grouped).map(group => {
    if (group.length === 1) {
      return group[0];
    }
    
    // Sort by updatedAt (newest first), fallback to createdAt
    return group.sort((a, b) => {
      const aTime = new Date(a.updatedAt || a.createdAt).getTime();
      const bTime = new Date(b.updatedAt || b.createdAt).getTime();
      return bTime - aTime; // Newest first
    })[0];
  });
}

// Helper function to get the latest version timestamp
export function getEntityTimestamp(entity: { createdAt: string; updatedAt: string }): Date {
  return new Date(entity.updatedAt || entity.createdAt);
}

// Filter out entities with empty or invalid names
export function filterValidEntities<T extends { name: string }>(entities: T[]): T[] {
  return entities.filter(entity => 
    entity.name && 
    entity.name.trim().length > 0 &&
    entity.name !== "here" && // Filter out test entries
    entity.name !== "london" // Filter out minimal test entries
  );
}

// Combined deduplication and filtering
// Deduplication utility for versioned entities
export function deduplicateByName<T extends Record<string, any>>(
  entities: T[]
): T[] {
  if (!entities || entities.length === 0) {
    return entities;
  }

  // Group entities by name/label (flexible)
  const grouped = entities.reduce((acc, entity) => {
    const name = entity.name || entity.label || entity.id;
    if (!name) return acc; // Skip entities without a name/label
    
    if (!acc[name]) {
      acc[name] = [];
    }
    acc[name].push(entity);
    return acc;
  }, {} as Record<string, T[]>);

  // For each group, keep only the latest version
  return Object.values(grouped).map(group => {
    if (group.length === 1) {
      return group[0];
    }
    
    // Sort by updatedAt/createdAt (newest first)
    return group.sort((a, b) => {
      const aTime = getEntityTimestamp(a);
      const bTime = getEntityTimestamp(b);
      return bTime.getTime() - aTime.getTime(); // Newest first
    })[0];
  });
}

// Helper function to get the latest version timestamp (flexible)
export function getEntityTimestamp(entity: Record<string, any>): Date {
  const updatedAt = entity.updatedAt || entity.updated_at;
  const createdAt = entity.createdAt || entity.created_at;
  
  if (updatedAt) {
    return new Date(updatedAt);
  }
  if (createdAt) {
    return new Date(createdAt);
  }
  
  // Fallback to current time if no timestamp
  return new Date(0); // Epoch time as fallback
}

// Filter out entities with empty or invalid names
export function filterValidEntities<T extends Record<string, any>>(entities: T[]): T[] {
  return entities.filter(entity => {
    const name = entity.name || entity.label;
    return name && 
           name.trim().length > 0 &&
           name !== "here" && // Filter out test entries
           name !== "london" && // Filter out minimal test entries  
           entity.description && entity.description.trim().length > 1; // Filter out minimal descriptions
  });
}

// Combined deduplication and filtering
export function cleanAndDeduplicateEntities<T extends Record<string, any>>(
  entities: T[]
): T[] {
  console.log('[Deduplication] Input entities:', entities.length);
  const filtered = filterValidEntities(entities);
  console.log('[Deduplication] After filtering:', filtered.length);
  const deduplicated = deduplicateByName(filtered);
  console.log('[Deduplication] After deduplication:', deduplicated.length);
  return deduplicated;
}
