# UI Refresh Functionality

## Overview

The AI Novel Prompter now includes comprehensive UI refresh functionality that allows users to reload data from the backend on-demand. This is especially useful when:

- Data has been modified externally (files changed outside the app)
- Multiple instances of the app are running
- Network/storage issues cause stale data
- Users want to ensure they're working with the latest data

## Features

### Manual Refresh Buttons

**Individual Refresh Buttons**: Each selector component (Characters, Locations, Codex, Rules, Task Types, Sample Chapters) has its own refresh button (🔄) that appears next to the edit button.

**Global Refresh Button**: Located in the header next to the storage indicator, refreshes all data types at once.

### Refresh Context System

- **DataRefreshProvider**: Manages global refresh state and functions
- **useDataRefresh**: Hook for accessing refresh functionality
- **Centralized State**: Tracks loading states and last refresh times
- **Error Handling**: Graceful error handling with user feedback

### Visual Feedback

- **Loading Animations**: Refresh icons spin during refresh operations
- **State Management**: Prevents multiple simultaneous refreshes
- **Error Indicators**: Visual feedback when refresh operations fail

## Technical Implementation

### Architecture

```
DataRefreshProvider (Context)
├── useDataRefresh (Hook)
├── RefreshButton (Component)
└── Individual Selector Components
```

### Key Files

- `frontend/src/contexts/DataRefreshContext.tsx` - Global refresh context
- `frontend/src/components/RefreshButton.tsx` - Reusable refresh button
- `frontend/src/hooks/useOptionManagement.ts` - Enhanced with refresh capability
- Updated selector components with refresh buttons

### Usage Example

```typescript
// Access refresh functions in any component
const { refreshFunctions, isRefreshing } = useDataRefresh();

// Refresh specific data type
await refreshFunctions.refreshCharacters();

// Refresh all data
await refreshFunctions.refreshAll();
```

## Configuration

The refresh behavior can be configured in `frontend/src/utils/refreshConfig.ts`:

- Auto-refresh intervals
- Debounce settings
- Visual feedback options
- Focus-based refresh triggers

## User Experience

### When to Use Refresh

1. **After External Changes**: When files are modified outside the application
2. **Multi-Instance Usage**: When running multiple app instances simultaneously
3. **Data Sync Issues**: When UI appears out of sync with backend storage
4. **Regular Updates**: Periodic refresh for latest data during long sessions

### Refresh Behavior

- **Non-Disruptive**: Current selections are preserved during refresh
- **Loading States**: Visual feedback during refresh operations
- **Error Recovery**: Graceful handling of failed refresh attempts
- **Debounced**: Prevents rapid successive refresh calls

## Future Enhancements

Potential improvements for future versions:

- Auto-refresh on window focus
- Refresh on dropdown open
- Background refresh with notifications
- Selective refresh (only changed entities)
- Refresh scheduling and intervals
- Keyboard shortcuts for refresh operations

## Troubleshooting

### Common Issues

1. **Refresh Button Not Working**
   - Check if Wails context is ready
   - Verify backend connectivity
   - Check browser console for errors

2. **Slow Refresh Performance**
   - May indicate large data sets
   - Backend storage optimization needed
   - Consider selective refresh

3. **Data Not Updating**
   - Verify backend file write permissions
   - Check data directory accessibility
   - Ensure proper error handling
