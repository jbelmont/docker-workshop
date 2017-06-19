import Vue from 'vue'
import VueResource from 'vue-resource'
import App from './components/Index.vue'
import router from './router'

Vue.config.productionTip = false

Vue.use(VueResource)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: {
    App
  }
})
