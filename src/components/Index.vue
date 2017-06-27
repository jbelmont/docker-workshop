<template>
  <div class="docker-container">
    <div class="navbar-component">
      <!-- Class `area` is a container -->
      <div class="navbar area">
        <!-- Logo -->
        <a href="#" class="dashboard">{{DASHBOARD}}</a>
        <!-- List of links -->
        <nav role="navigation" id="navigation" class="list">
          <a href="#" class="item -link">{{USERS}}</a>
          <a href="#" class="item -link">{{SETTINGS}}</a>
          <a href="#" class="item -link">{{PROFILE}}</a>
          <span class="item"><i class="fa fa-search"></i></span>
        </nav>
        <!-- Button to toggle the display menu  -->
        <button data-collapse data-target="#navigation" class="toggle">
          <!-- Hamburger icon -->
          <span class="icon"></span>
        </button>
      </div>
    </div>
    <div class="users-area">
      <table class="users-area-table">
        <thead>
          <tr>
            <td>{{ID}}</td>
            <td>{{FIRST_NAME}}</td>
            <td>{{LAST_NAME}}</td>
            <td>{{GENDER}}</td>
            <td>{{EMAIL}}</td>
            <td>{{REMOVE_USER}}</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in users"
              v-bind:key="user.id">
            <td>{{ user.id }}</td>
            <td>{{ user.first_name }}</td>
            <td>{{ user.last_name }}</td>
            <td>{{ user.gender }}</td>
            <td>{{ user.email }}</td>
            <td><button :data-id=user.id v-on:click="removeUser">{{REMOVE_USER}}</button></td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import * as constants from '../constants'
const {
  DASHBOARD,
  USERS,
  SETTINGS,
  PROFILE,
  ID,
  FIRST_NAME,
  LAST_NAME,
  GENDER,
  EMAIL,
  REMOVE_USER
} = constants

export default {
  name: 'index',
  data () {
    return {
      users: [],
      DASHBOARD,
      USERS,
      SETTINGS,
      PROFILE,
      ID,
      FIRST_NAME,
      LAST_NAME,
      GENDER,
      EMAIL,
      REMOVE_USER
    }
  },
  mounted: function () {
    return this.getUsers()
  },
  methods: {
    getUsers: function () {
      this.$http.get('api/v1/users').then(function (data) {
        this.users = data.body
      }, function (err) {
        console.error(err)
      })
    },
    removeUser: function (event) {
      console.log(event)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import '../scss/variables';
// Scaffolding
*, *:before, *:after {
  box-sizing: border-box;
}

body {
  background-color: #f5f5f5;
  color: #333;
  font-size: 14px;
  font-family: Verdana, Arial, sans-serif;
  line-height: 20px;
}
a {
  text-decoration: none; transition: all 0.3s linear 0s;
}

.users-area {
  display: flex;
  flex-direction: row;

  &-table {
    width: 100%;
  }
}

.area {
  display: flex; flex-flow: row wrap; align-items: stretch; margin-left: auto; margin-right: auto;
  @media (min-width: 768px) { width: 750px; }
  @media (min-width: 992px) { width: 970px; }
  @media (min-width: 1200px) { width: 1140px; }
}

// Navigation component
// ----------

// Component skeleton
.navbar-component {
  background-color: $navbar-background;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.16), 0 2px 10px rgba(0, 0, 0, 0.12);

  & > .navbar {
    justify-content: space-between;
  }
}

td, th {
  text-align: center;
  border: solid 1px #ccc;
  font-size: 22px;
  padding: 6px;
  word-wrap: break-word;
}

.navbar {
  & > .dashboard {
    display: block;
    font-size: 16px;
    color: #777;
    margin: round(($navbar-height - 20) / 2);
  }

  // Toggle button
  & > .toggle {
    border: 0;
    background-color: transparent;
    outline: none;
    border: 0;
    display: inline-block;
    background-color: transparent;
    background-image: none;
    vertical-align: middle;
    text-align: center;
    white-space: nowrap;
    cursor: pointer;
    touch-action: manipulation;
    user-select: none;
    padding: round(($navbar-height - 20) / 2);

     @media (min-width: $navbar-collapse-breakpoint) {
       display: none;
     }
  }

  & > .toggle > .icon {
    position: relative;
    margin-top: 8px;
    margin-bottom: 8px;

    &, &:before, &:after {
      display: block;
      width: 24px;
      height: 3px;
      transition: background-color 0.3s linear, transform 0.3s linear;
      background-color: #555555;
    }

    &:before, &:after {
      position: absolute; content: "";
    }
    &:before {
      top: -8px;
    }
    &:after {
      top: 8px;
    }
  }

  & > .toggle.-active > .icon {
    background-color: transparent;

    &:before { transform: translateY(8px) rotate(45deg); }
    &:after { transform: translateY(-8px) rotate(-45deg); }
  }

  // List of items
  & > .list {
    display: none;
    flex-flow: row nowrap;
    align-items: center;
    white-space: nowrap;

    @media (min-width: $navbar-collapse-breakpoint) {
      display: flex;
    }

    @media (max-width: $navbar-collapse-breakpoint) {
      position: fixed;
      top: $navbar-height;
      left: 0;
      width: 100%;
      overflow-y: hidden;
      overflow-x: auto;
      border-top: 1px solid $navbar-border;
      background-color: $navbar-background;
    }

    &.-on {
      display: flex;
    }
  }

  & > .list > .item {
    display: block;
    flex-shrink: 0;
    height: $navbar-height;
    line-height: $navbar-height;
    padding-left: round(($navbar-height - 20) / 2);
    padding-right: round(($navbar-height - 20) / 2);
    text-transform: uppercase;
    color: $navbar-item-color;
    font-size: $navbar-item-font-size;
  }

  & > .list > .item.-link {
    line-height: $navbar-height + $navbar-item-border-width;
    color: $navbar-item-color;
    border-bottom: $navbar-item-border-width solid $navbar-item-border;

    &.-active, &:hover, &:focus {
      color: $navbar-item-active-color;
      border-bottom-color: $navbar-item-active-border;
    }
  }
}
</style>
