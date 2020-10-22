import { createPlugin } from '@backstage/core';
import Payment from './components/Payment';
import Login from './components/Login';
import Table from './components/Table';
import Home from './components/Home';
export const plugin = createPlugin({
  id: 'payment',
  register({ router }) {
    router.registerRoute('/payment', Payment);
    router.registerRoute('/home', Home);
    router.registerRoute('/table', Table);
    router.registerRoute('/', Login);
  },
});
