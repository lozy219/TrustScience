//
//  LaunchManager.swift
//  TrustScienceCore
//
//  Created by Wang Jinghan on 10/3/19.
//  Copyright © 2019 TrustScience. All rights reserved.
//

import Foundation

public class LaunchManager: Manager {
    
    public override init() {
        super.init()
    }
    
    public func launchType() -> LaunchType {
        return .main
    }
    
}

public enum LaunchType {
    case main
}

