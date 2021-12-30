<template>
  <div class="position-absolute top-0 end-0">
      <button type="button" class="btn btn-secondary btn-sm pb-0" id="dropdownMenuButton1" data-bs-toggle="dropdown" aria-expanded="false">
        <p class="icon"></p>
        <p class="icon"></p>
        <p class="icon"></p> 
      </button>
      <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton1">
        <li><a class="dropdown-item" @click="sendDelete">delete</a></li>
      </ul>
  </div>
</template>

<script>
export default {
  name: "Editor",
  props: {
    target: String,
    id: Number,
  },
  setup(props, context) {
    const sendDelete = () => {
      context.emit('sendDelete', props.target, props.id)
    }
    return {
      sendDelete
    }
  }
}
</script>
<style scoped>
.icon {
  display: inline-block;
  width: 4px;
  height: 4px;
  margin: 4px;
  border-radius: 50%;
  background-color: white;
}
</style>
