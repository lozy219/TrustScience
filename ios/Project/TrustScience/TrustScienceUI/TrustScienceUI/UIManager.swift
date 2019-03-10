//
//  UIManager.swift
//  TrustScience
//
//  Created by Wang Jinghan on 10/3/19.
//  Copyright © 2019 TrustScience. All rights reserved.
//

import UIKit
import TrustScienceCore

public class UIManager: Manager {
    
    private let window: UIWindow
    
    init(window: UIWindow) {
        self.window = window
        
        super.init()
    }
    
    func launch(type: LaunchType) {
        switch type {
        case .main:
            launchMainScreen()
        }
    }
    
    func frontmostViewController() -> UIViewController? {
        var frontmostVC: UIViewController? = window.rootViewController
        while let presentedVC = frontmostVC?.presentedViewController {
            frontmostVC = presentedVC
        }
        return frontmostVC
    }
    
}

extension UIManager {
    private func launchMainScreen() {
        let vc = ClientStream.session.factoryUI.buildMainViewController()
        replaceRootViewController(vc)
    }
    
    private func replaceRootViewController(_ viewController: ViewController) {
        if let oldViewController = window.rootViewController {
            
            // Replacing rootVC does not dismiss any presented VC causing weird UI and leak
            oldViewController.dismissIfNeeded(animated: false) { [weak self] in
                guard let me = self else { return }
                
                let transition = CATransition()
                transition.duration = 0.3
                transition.type = .fade
                transition.subtype = .fromBottom
                transition.timingFunction = CAMediaTimingFunction(name: .easeOut)
                me.window.layer.add(transition, forKey: kCATransition)
                
                me.window.rootViewController = viewController
                me.window.makeKeyAndVisible()
            }
        } else {
            window.rootViewController = viewController
            window.makeKeyAndVisible()
        }
    }
    
}

// MARK: - UIViewController
private extension UIViewController {
    func dismissIfNeeded(animated: Bool, completion: EmptyBlock?) {
        
        guard let pvc = presentedViewController else {
            completion?()
            return
        }
        
        guard !pvc.isBeingDismissed else {
            // Shouldn't happen, otherwise will need to wait.
            // If we do something while dismissing -> can cause issues.
            // Eg: changing window rootViewController -> cause leak (old view not be removed from window)
            DispatchQueue.main.asyncAfter(deadline: .now() + 0.1) {[weak self] in
                self?.dismissIfNeeded(animated: animated, completion: completion)
            }
            return
        }
        
        dismiss(animated: animated) {
            DispatchQueue.main.async { completion?() }
        }
    }
}

