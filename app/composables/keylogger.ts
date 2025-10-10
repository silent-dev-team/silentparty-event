
type Options = {
  immediate?: boolean;
  timeout?: number;
};

import { onUnmounted, readonly, ref } from 'vue';

const paused = ref(false);

const commandsPrefix = ['%']
export const enter = '%enter;'

export const useKeyLogger = (callback?: (logs:string) => void, opts?: Options) => {
  const pb = usePocketbase();
  const keys = ref<string>();
  let timeoutId: ReturnType<typeof setTimeout>;

  const logKey = (event: KeyboardEvent) => {
    if (commandsPrefix.includes(event.key)) {
      resume();
    }
    if (paused.value) return;
    event.preventDefault();
    console.log(`Key pressed: ${event.key} - Code: ${keys.value}`);
    if (event.key === 'Backspace') {
      keys.value = keys.value ? keys.value.slice(0, -1) : '';
      return;
    }
    if (event.key.length > 1) {
      // Ignore special keys like Shift, Ctrl, etc.
      if (['Enter', 'Tab', 'Space'].includes(event.key)) {
        if (callback && keys.value) {
          callback(keys.value);
        }
        keys.value = ''
      }
      return;
    }
    keys.value = keys.value ? keys.value + event.key : event.key;
    if (opts?.timeout) {
      clearTimeout(timeoutId);
      timeoutId = setTimeout(() => {
        if (callback && keys.value) {
          callback(keys.value);
        }
        keys.value = '';
      }, opts.timeout);
    }
  };

  const startLogging = () => {
    console.log('Key logging started');
    window.addEventListener('keydown', logKey);
  };

  const stopLogging = () => {
    console.log('Key logging stopped');
    window.removeEventListener('keydown', logKey);
  };

  const pause = () => paused.value = true;
  const resume = () => paused.value = false;

  const reset = () => {
    keys.value = '';
  };

  onMounted(() => {
    if (opts?.immediate && pb.authStore.isValid) {
      startLogging();
    }
  });

  onUnmounted(() => {
    stopLogging();
  });

  return {
    startLogging,
    stopLogging,
    logs: readonly(keys),
    reset,
    pause,
    resume,
    paused,
  };
};
