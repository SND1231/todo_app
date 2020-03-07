// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import axios from 'axios'
import App from './App'

Vue.config.productionTip = false

const baseUrl = 'http://192.168.33.10:8080/'

// todoリスト
new Vue({
  el: '#todo_list',
  data: function () {
    return {
      todos: null
    }
  },
  mounted () {
    axios
      .get(baseUrl + 'todo')
      .then(response => (this.todos = response.data))
      .catch(error => console.log(error))
  },
  methods: {
    createTodo: function (event, value) {
      var comment = this.$refs.comment
      // 入力がなければ何もしないで return
      if (!comment.value.length) {
        return
      }
      axios
        .post(baseUrl + 'todo', {Content: comment.value})
        .catch(error => console.log(error))
        .finally(() => location.reload())
    },
    editTodo: function (id) {
      var contents = this.$refs.content
      var edit_no = -1
      for (var i = 0; i < contents.length; i++) {
        if(id == this.$refs.id[i].textContent){
          edit_no = i
          break
        } 
      }
      console.log(edit_no)
      if (edit_no==-1){
        return
        location.reload()
      }
      var content = contents[edit_no]
      // 入力がなければ何もしないで return
      if (!content.value.length) {
        location.reload()
      }
      axios
        .put(baseUrl + 'todo/' + String(id), {Content: content.value})
        .catch(error => console.log(error))
        .finally(() => location.reload())
    },
    deleteTodo: function (id) {
      axios
        .delete(baseUrl + 'todo/' + String(id))
        .catch(error => console.log(error))
        .finally(() => location.reload())
    },
  }
})
