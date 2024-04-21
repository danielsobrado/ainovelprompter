// assets
import { IconLayersLinked, IconHelp } from '@tabler/icons-react';

// constant
const icons = { IconLayersLinked, IconHelp };

// ==============================|| SAMPLE PAGE & DOCUMENTATION MENU ITEMS ||============================== //

const other = {
  id: 'sample-docs-roadmap',
  type: 'group',
  children: [
    {
      id: 'generate-prompt',
      title: 'Generate Prompt',
      type: 'item',
      url: '/generate-prompt',
      icon: icons.IconLayersLinked,
      breadcrumbs: false
    },
    {
      id: 'drusniel',
      title: 'Drusniel',
      type: 'item',
      url: 'https://drusniel.com/',
      icon: icons.IconHelp,
      external: true,
      target: true
    }
  ]
};

export default other;
