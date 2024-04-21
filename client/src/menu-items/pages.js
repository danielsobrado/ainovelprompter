// assets
import { IconWriting, IconBrandHipchat } from '@tabler/icons-react';

// constant
const icons = {
  IconWriting,
  IconBrandHipchat
};

// ==============================|| EXTRA PAGES MENU ITEMS ||============================== //

const pages = {
  id: 'localModels',
  title: 'Local Models',
  caption: 'Fine Tuning & Chat',
  type: 'group',
  children: [
    {
      id: 'fineTuning',
      title: 'Fine Tuning',
      type: 'collapse',
      icon: icons.IconWriting,

      children: [
        {
          id: 'login3',
          title: 'Login',
          type: 'item',
          url: '/pages/login/login3',
          target: true
        },
        {
          id: 'register3',
          title: 'Register',
          type: 'item',
          url: '/pages/register/register3',
          target: true
        }
      ]
    },
    {
      id: 'chat',
      title: 'Ollama Chat',
      type: 'item',
      url: '/pages/ollama-chat/OllamaChat',
      icon: icons.IconBrandHipchat,
      breadcrumbs: false
    }
  ]
};

export default pages;
