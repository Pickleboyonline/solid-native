import { NativeModulesProxy, EventEmitter, Subscription } from 'expo-modules-core';

// Import the native module. On web, it will be resolved to TestExpoModule.web.ts
// and on native platforms to TestExpoModule.ts
import TestExpoModule from './TestExpoModule';
import TestExpoModuleView from './TestExpoModuleView';
import { ChangeEventPayload, TestExpoModuleViewProps } from './TestExpoModule.types';

// Get the native constant value.
export const PI = TestExpoModule.PI;

export function hello(): string {
  return TestExpoModule.hello();
}

export async function setValueAsync(value: string) {
  return await TestExpoModule.setValueAsync(value);
}

const emitter = new EventEmitter(TestExpoModule ?? NativeModulesProxy.TestExpoModule);

export function addChangeListener(listener: (event: ChangeEventPayload) => void): Subscription {
  return emitter.addListener<ChangeEventPayload>('onChange', listener);
}

export { TestExpoModuleView, TestExpoModuleViewProps, ChangeEventPayload };
