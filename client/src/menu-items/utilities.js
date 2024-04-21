// assets
import { IconDatabase } from '@tabler/icons-react';

// constant
const icons = {
  IconDatabase
};

// ==============================|| UTILITIES MENU ITEMS ||============================== //

const utilities = {
  id: 'Prompt',
  title: 'Prompt',
  type: 'group',
  children: [
    {
      id: 'trait-types',
      title: 'Trait Types',
      type: 'item',
      url: '/trait-types',
      icon: icons.IconDatabase,
      breadcrumbs: false
    },
    {
      id: 'texts',
      title: 'Texts',
      type: 'item',
      url: '/texts',
      icon: icons.IconDatabase,
      breadcrumbs: false
    },
    {
      id: 'standard-prompts',
      title: 'Standard Prompts',
      type: 'item',
      url: '/standard-prompts',
      icon: icons.IconDatabase,
      breadcrumbs: false
    },
    {
      id: 'chapters',
      title: 'Chapters',
      type: 'collapse',
      icon: icons.IconDatabase,
      children: [
        {
          id: 'table-chapters',
          title: 'Chapters',
          type: 'item',
          url: '/model/table-chapters',
          breadcrumbs: false
        },
        {
          id: 'table-chapter-details',
          title: 'Chapter Details',
          type: 'item',
          url: '/model/table-chapter-details',
          breadcrumbs: false
        }
      ]
    }
  ]
};

export default utilities;
