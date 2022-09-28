// Copyright 2022 The Kubermatic Kubernetes Platform contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import {KeyValueEntry} from '../types/common';

export const CLUSTER_DEFAULT_NODE_SELECTOR_NAMESPACE = 'clusterDefaultNodeSelector';
export const CLUSTER_DEFAULT_NODE_SELECTOR_TOOLTIP =
  'Namespace clusterDefaultNodeSelector requires nodes with the labels defined in the Label Selector to start the cluster. The labels are enforced. Edit the label in the Label Selector.';
export const CLUSTER_DEFAULT_NODE_SELECTOR_HINT =
  'Namespace <b>clusterDefaultNodeSelector</b> requires nodes with the labels defined in the Label Selector to start the cluster. The labels are enforced.';

export function handleClusterDefaultNodeSelector(
  labels: Record<string, string>,
  nodeConfig: Record<string, string>,
  previousClusterDefaultNodeSelectorNamespace: KeyValueEntry,
  onComplete: (entry: KeyValueEntry, labels: Record<string, string>) => void
): void {
  const [prevKey] = previousClusterDefaultNodeSelectorNamespace ?? [];

  if (Object.hasOwnProperty.call(labels, prevKey)) {
    delete labels[prevKey];
  }

  const [currKey, currValue] = nodeConfig[CLUSTER_DEFAULT_NODE_SELECTOR_NAMESPACE]?.split('=') ?? [];

  onComplete(currKey && currValue ? [currKey, currValue] : null, labels);
}
