import { createPlugin } from '@backstage/core';
import Login from './components/Login'
import Area from './components/Area'


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/area', Area);
  },
});
