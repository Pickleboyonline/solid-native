import { requireNativeViewManager } from 'expo-modules-core';
import * as React from 'react';

import { TestExpoModuleViewProps } from './TestExpoModule.types';

const NativeView: React.ComponentType<TestExpoModuleViewProps> =
  requireNativeViewManager('TestExpoModule');

export default function TestExpoModuleView(props: TestExpoModuleViewProps) {
  return <NativeView {...props} />;
}
