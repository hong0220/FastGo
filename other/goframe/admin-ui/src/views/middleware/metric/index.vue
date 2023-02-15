<template>
  <iframe
      src="http://127.0.0.1:9090/metric"
      height="100%"
      width="100%"
      id="iframe"
      @load="adjustIframe"
  />
</template>

<script lang="ts">
import {defineComponent} from 'vue';

export default defineComponent({
  name: 'middlewareMetric',
  methods: {
    adjustIframe() {
      var ifm = document.getElementById("iframe");
      ifm.height = document.documentElement.clientHeight;
      ifm.width = document.documentElement.clientWidth;
    }
  }
});
</script>
