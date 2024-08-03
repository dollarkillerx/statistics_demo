import { RouteRecordRaw } from 'vue-router';
import { Layout } from '@/router/constant';
import { ProjectOutlined } from "@vicons/antd";
import { renderIcon } from '@/utils/index';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/task',
    name: 'Task',
    component: Layout,
    meta: {
      sort: 1,
      icon: renderIcon(ProjectOutlined),
    },
    children: [
      {
        path: 'task_index',
        name: `task_index`,
        meta: {
          title: '运行任务',
        },
        component: () => import('@/views/task/index/index.vue'),
      },
    ],
  },
];

export default routes;
