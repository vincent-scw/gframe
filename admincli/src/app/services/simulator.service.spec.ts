import { TestBed } from '@angular/core/testing';

import { SimulatorService } from './simulator.service';

describe('SimulatorService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: SimulatorService = TestBed.get(SimulatorService);
    expect(service).toBeTruthy();
  });
});
