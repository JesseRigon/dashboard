// Copyright 2020 The Kubermatic Kubernetes Platform contributors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import {HttpClientModule} from '@angular/common/http';
import {async, ComponentFixture, TestBed} from '@angular/core/testing';
import {BrowserModule} from '@angular/platform-browser';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {RouterTestingModule} from '@angular/router/testing';
import {AppConfigService} from '../../../app-config.service';
import {ApiService, DatacenterService, ParamsService, PresetsService, ProjectService} from '../../../core/services';
import {NODE_DATA_CONFIG, NodeDataMode} from '../../../node-data/config';
import {NodeDataService} from '../../../node-data/service/service';
import {ClusterService} from '../../../shared/services/cluster.service';
import {ApiMockService} from '../../../testing/services/api-mock.service';
import {ProjectMockService} from '../../../testing/services/project-mock.service';
import {WizardService} from '../../../wizard/service/wizard';
import {SharedModule} from '../../shared.module';
import {MachineNetworkComponent} from './component';
import {DatacenterMockService} from '../../../testing/services/datacenter-mock.service';

const modules: any[] = [BrowserModule, HttpClientModule, BrowserAnimationsModule, RouterTestingModule, SharedModule];

describe('MachineNetworksComponent', () => {
  let component: MachineNetworkComponent;
  let fixture: ComponentFixture<MachineNetworkComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [...modules],
      providers: [
        AppConfigService,
        ClusterService,
        NodeDataService,
        ParamsService,
        PresetsService,
        WizardService,
        {provide: ProjectService, useValue: ProjectMockService},
        {provide: ApiService, useValue: ApiMockService},
        {provide: NODE_DATA_CONFIG, useValue: NodeDataMode.Wizard},
        {provide: DatacenterService, useClass: DatacenterMockService},
      ],
    }).compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MachineNetworkComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create the machine network component', () => {
    expect(component).toBeTruthy();
  });

  it('expecting form to be valid', () => {
    const machineNetworks = component._networkArray;
    expect(machineNetworks.valid).toBeTruthy();
  });

  it('expecting form to be invalid', () => {
    component.add();
    const machineNetworks = component._networkArray;
    machineNetworks.controls[0].setValue({
      cidr: '192.182.0.0',
      dnsServers: ['8.8.8.8'],
      gateway: '192.180.0.2',
    });
    expect(machineNetworks.controls[0].valid).toBeFalsy();
  });
});
