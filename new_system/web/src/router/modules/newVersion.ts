import { RouteRecordRaw } from 'vue-router';
import { Layout } from '@/router/constant';
import { SketchOutlined } from '@vicons/antd';
import { renderIcon, renderNew } from '@/utils/index';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/newversion',
    name: 'https://www.naiveadmin.com',
    component: Layout,
    meta: {
      title: 'Pro 版本',
      icon: renderIcon(SketchOutlined),
      sort: 1,
    },
  },
];

export default routes;
