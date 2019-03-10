//
//  UIFactory.swift
//  TrustScience
//
//  Created by Wang Jinghan on 10/3/19.
//  Copyright © 2019 TrustScience. All rights reserved.
//

import Foundation
import TrustScienceUI

class UIFactory {
    struct Dependencies {
        let managerUI: UIManager
    }
    
    let deps: Dependencies
    init(deps: Dependencies) {
        self.deps = deps
    }
    
    func buildMainViewController() -> ViewController {
        return ViewController()
    }
}
