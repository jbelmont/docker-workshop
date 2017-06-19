import Vue from 'vue'
import Router from 'vue-router'
import Main from '../components/Index.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Main
    }
  ]
})
