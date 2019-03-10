//
//  UIFactory.swift
//  TrustScience
//
//  Created by Wang Jinghan on 10/3/19.
//  Copyright © 2019 TrustScience. All rights reserved.
//

import Foundation
import TrustScienceCore

public class UIFactory {
    struct Dependencies {
        let managerUI: UIManager
        let managerHTTP: HTTPManager
    }
    
    let deps: Dependencies
    init(deps: Dependencies) {
        self.deps = deps
    }
    
    func buildMainViewController() -> ViewController {
        return ViewController(factoryUI: self)
    }
}
