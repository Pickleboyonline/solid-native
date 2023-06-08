import * as React from 'react';

import { TestExpoModuleViewProps } from './TestExpoModule.types';

export default function TestExpoModuleView(props: TestExpoModuleViewProps) {
  return (
    <div>
      <span>{props.name}</span>
    </div>
  );
}
