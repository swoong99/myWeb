import { createRouter, createWebHistory } from 'vue-router';
import Introduction from '@/components/Introduction.vue';
import Projects from '@/components/Projects.vue';
import Research from '@/components/Research.vue';

const routes = [
    { path: '/', redirect: '/introduction' },
    { path: '/introduction', component: Introduction },
    { path: '/projects', component: Projects },
    { path: '/research', component: Research },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;